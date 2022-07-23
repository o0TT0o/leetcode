
var MyQueue = function() {
    this.stack = []
    this.stack2 = []
};

/** 
 * @param {number} x
 * @return {void}
 */
MyQueue.prototype.push = function(x) {
    this.stack.push(x)
};

/**
 * @return {number}
 */
MyQueue.prototype.pop = function() {
    if(this.stack2.length===0){
        while(this.stack.length>0){
            this.stack2.push(this.stack.pop())
        }
    }
    return this.stack2.pop()
};
 
/**
 * @return {number}
 */
MyQueue.prototype.peek = function() {
    if(this.stack2.length<=0){
        while(this.stack.length>0){
            this.stack2.push(this.stack.pop())
        }
    }
    return this.stack2[this.stack2.length-1]
};

/**
 * @return {boolean}
 */
MyQueue.prototype.empty = function() {
    if(this.stack.length<=0 && this.stack2.length<=0){
        return true
    }else{
        return false
    }
};

/** 
 * Your MyQueue object will be instantiated and called as such:
 * var obj = new MyQueue()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.peek()
 * var param_4 = obj.empty()
 */