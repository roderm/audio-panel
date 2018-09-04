let websocketInstances = {};

class EventEmitter {
    constructor() {
        this.events = {};
        this.oneTimeEvent = {}
        this.subscribe = this.subscribe.bind(this)
        this.emit = this.emit.bind(this)
        this.once = this.once.bind(this)
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
                var i = 1;
                for(var i = 0; i < event.length; i++){
                    event[i].call(null, ...data);
                }
            }
        }

        const event = this.events[eventName];
        const once = this.oneTimeEvent[eventName];
        emitSubs(event);
        emitSubs(once);
        this.oneTimeEvent[eventName] = {};
    }

    once(eventName, fn) {
        if (!this.oneTimeEvent[eventName]) {
            this.oneTimeEvent[eventName] = [];
        }

        this.oneTimeEvent[eventName].push(fn);

        return () => {
            this.oneTimeEvent[eventName] = this.oneTimeEvent[eventName].filter(eventFn => fn !== eventFn);
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
            this.setUpSocket = this.setUpSocket.bind(this)
            this.getSocket = this.getSocket.bind(this)
            this.rx = this.rx.bind(this)
            this.sub = this.sub.bind(this)
            api.SocketEmitter = new EventEmitter()
            api.socket = new WebSocket(_path)
            api.socket.addEventListener("open", () => {
                api.socketIsConnected = true;
                api.SocketEmitter.emit("open")
                api.SocketEmitter.emit("connected")
                api.setUpSocket()
            })
        }
        return websocketInstances[_path]
    }
    setUpSocket() {
        let api = this;
        let reconnTimeout
        let reconnFn = () => {
            clearTimeout(reconnTimeout)
            reconnTimeout = setTimeout(() => {
                api.socketIsConnected = false;
                let ws = new WebSocket(api.ws_path)
                ws.addEventListener("open", () => {
                    api.socket = ws
                    api.setUpSocket()
                    api.socketIsConnected = true;
                    api.SocketEmitter.emit("reconnected")
                    api.SocketEmitter.emit("connected")
                })
                ws.addEventListener("error", () => {
                    reconnFn()
                })
            }, 5000)
        }

        api.socket.addEventListener("close", reconnFn)
        api.socket.addEventListener("message", function (e) {
            let data = JSON.parse(e.data)
            api.SocketEmitter.emit(data.id, data.result, data.error)
        })
    }
    getSocket(cb) {
        let promise = new Promise(function(resolve, reject){
            if (this.socketIsConnected) {
                resolve(this.socket)
            }
            this.SocketEmitter.once("connected", function() {
                resolve(this.socket)
            }.bind(this))
        }.bind(this))
        return promise
    }
    wsSub() {
        return this.SocketEmitter
    }
    rx(action, data) {
        let req = {
            method: action,
            id: this.createId(),
            params: data
        }
        let promise = new Promise((resolve, reject) => {
            this.getSocket().then((ws) => {
                this.SocketEmitter.once(req.id, (result, error) => {
                    if (error) {
                        if (error.code != 0) {
                            reject(error)
                        }
                    }
                    resolve(result)
                })
                ws.send(JSON.stringify(req))
            })
        })
        return promise
    }
    sub(action, data) {
        let req = {
            method: action,
            id: this.createId(),
            params: data
        }

        let subscription = new EventEmitter()
        this.getSocket().then((ws) => {
            this.SocketEmitter.subscribe(req.id, (...data) => {
                subscription.emit("on" + action, ...data)
            })
            ws.send(JSON.stringify(req))
        })
        return subscription
    }
}

export default WebsocketApi;