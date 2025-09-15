let promise = new Promise ((resolve,reject)=>setTimeout(()=>reject("Promise not fullfilled"))
.then(reject=>console.log("a:");


)