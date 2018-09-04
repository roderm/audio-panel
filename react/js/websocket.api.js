let websocketInstances = {};

class EventEmitter {
    constructor() {
        this.events = {};
        this.onTimeEvent = {}
    }
    subscribe(eventName, fn) {
        if (!this.events[eventName]) {
            this.events[eventName] = [];
        }

        this.events[eventName].push(fn);

        return () => {
            this.events[eventName] = this.events[eventName].filter(eventFn => fn !== eventFn);
        }
    }

    emit(eventName, ...data) {
        let emitSubs = (event) => {
            if (event && event.constructor === Array) {
                event.every(fn => {
                    fn.call(null, ...data);
                });
            }
        }

        const event = this.events[eventName];
        const once = this.onTimeEvent[eventName];
        emitSubs(event);
        emitSubs(once);
        this.onTimeEvent[eventName] = {};
    }

    once(eventName, fn) {
        if (!this.onTimeEvent[eventName]) {
            this.onTimeEvent[eventName] = [];
        }

        this.onTimeEvent[eventName].push(fn);

        return () => {
            this.onTimeEvent[eventName] = this.onTimeEvent[eventName].filter(eventFn => fn !== eventFn);
        }
    }
}

class WebsocketApi {
    constructor(_path) {
        if (!websocketInstances[_path]) {
            websocketInstances[_path] = this

            let api = this
            api.ws_path = _path
            api.idcounter = 0;
            api.socketIsConnected = false;
            api.createId = () => {
                return (api.idcounter++).toString();
            }
            let socket = new WebSocket(_path)
            api.setUpSocket(socket)
            api.SocketEmitter = new EventEmitter()
        }
        return websocketInstances[_path]
    }
    setUpSocket(socket){
        let api = this;
        let reconnFn = () => {
            api.socketIsConnected = false;
            // TODO: add timeout
            let ws = new WebSocket(api.ws_path)
            api.setUpSocket(ws)
        }
        socket.addEventListener("open", () => {
            api.socketIsConnected = true;
        })
        socket.addEventListener("close", reconnFn)
        socket.addEventListener("error", reconnFn)
        socket.addEventListener("message", function (e) {
            let data = JSON.parse(e.data)
            api.SocketEmitter.emit(data.id, data.result, data.error)
        })
        api.socket = socket
    }
    getSocket() {
        let api = this
        let promise = new Promise((resolve, reject) => {
            if (this.socketIsConnected) {
                resolve(this.socket)
                return
            }
            this.socket.addEventListener("open", () => {
                resolve(this.socket)
            })
        })
        return promise
    }
    rx(action, data) {
        let api = this;
        let req = {
            method: action,
            id: this.createId(),
            params: data
        }
        let promise = new Promise((resolve, reject) => {
            api.getSocket().then((ws) => {
                ws.send(JSON.stringify(req))
                api.SocketEmitter.once(req.id, (...data) => {
                    resolve(...data)
                })
            })
        })
        return promise
    }
    sub(action, data) {
        let api = this;
        let req = {
            method: action,
            id: this.createId(),
            params: data
        }
        let subscribtion = new EventEmitter()
        api.getSocket().then((ws) => {
            ws.send(JSON.stringify(req))
            api.SocketEmitter.subscribe(req.id, (...data) => {
                subscribtion.emit("on" + action, ...data)
            })
        })
        return subscribtion
    }
}

export default WebsocketApi;