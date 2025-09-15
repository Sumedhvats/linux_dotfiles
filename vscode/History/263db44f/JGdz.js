// Object.prototype.print=()=>{
//     console.log("hwllo");
    
// }
// p={
//     "name":"sumedh"
// }
// p.print()

Array.prototype.sum=()=>{
    let sum=0;
    for (let i =0;i<this.length;i++){
        sum += this[i]
    }
    return sum
}
let a1=[1,2,3,4]
console.log(a1.sum());
