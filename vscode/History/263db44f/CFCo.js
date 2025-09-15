async function name(params) {
    let promise = new Promise((resolve, reject) => 
      setTimeout(() => reject("Promise not fulfilled"), 1000)
    );
    promise
      .catch(err => { console.log("a:", err); throw err; })
      .catch(err => { console.log("b:", err); throw err; })
      .catch(err => { console.log("c:", err); throw err; })
      .catch(err => { console.log("d:", err); });
    
}