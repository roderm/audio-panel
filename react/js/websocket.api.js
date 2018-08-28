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

    emit(eventName, data) {
        let emitSubs = (event) => {
            if (event && event.constructor === Array) {
                event.every(fn => {
                    fn.call(null, data);
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
            api.idcounter = 0;
            api.socketIsConnected = false;
            api.createId = () => {
                return (api.idcounter++).toString();
            }
            api.socket = new WebSocket(_path)
            api.socket.addEventListener("open", function () {
                api.socketIsConnected = true;
            })
            api.getSocket = () => {
                let promise = new Promise(function (resolve, reject) {
                    if (api.socketIsConnected) {
                        resolve(api.socket)
                        return
                    }
                    api.socket.addEventListener("open", function () {
                        resolve(api.socket)
                    })
                })
                return promise
            }
            api.SocketEmitter = new EventEmitter()
            api.socket.addEventListener("message", function (e) {
                let data = JSON.parse(e.data)
                console.debug(e)
                api.SocketEmitter.emit(data.id, data.result)
            })
        }
        return websocketInstances[_path]
    }
    rx(action, data) {
        let api = this;
        let req = {
            method: action,
            id: this.createId(),
            params: data
        }
        let promise = new Promise(function (resolve, reject) {
            api.getSocket().then(function (ws) {
                ws.send(JSON.stringify(req))
                api.SocketEmitter.once(req.id, function (data) {
                    resolve(data)
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
        api.getSocket().then(function (ws) {
            ws.send(JSON.stringify(req))
            api.SocketEmitter.subscribe(req.id, function (data) {
                subscribtion.emit("on" + action, data)
            })
        })
        return subscribtion
    }
}

export default WebsocketApi;