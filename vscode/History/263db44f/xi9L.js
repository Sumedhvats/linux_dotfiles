async function getUser() {
    const response = await fetch('https://jsonplaceholder.typicode.com/users/1');
    const user = await response.json();

    console.log("Name:", user.name);
    console.log("Email:", user.email);
    console.log("City:", user.address.city);
}

getUser();
