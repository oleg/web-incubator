export default class Util {
    static _bind(target, ...methods) {
        for (let m of methods) {
            target[m] = target[m].bind(target);
        }
    }

    static logDebug(message) {
        console.log(message)
    }

}