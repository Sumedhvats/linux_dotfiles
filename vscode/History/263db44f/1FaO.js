async function getUserData() {
    
}
fetch('https://jsonplaceholder.typicode.com/users/1')
  .then(response => response.json())
  .then(user => {
    console.log("Name:", user.name);
    console.log("Email:", user.email);
    console.log("City:", user.address.city);
  });
