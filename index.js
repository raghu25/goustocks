var user = {
    username : 'darshan',
    password : '123',
    isNew : false
}

function main(u) {
    u.isNew = true;
    u.username = "Raghu"
}

console.log("before passing" , user)
main(user)
console.log(user);

