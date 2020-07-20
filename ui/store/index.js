export const state = () => ({
    counter: 0,
    user: {
        isLoggedIn: false,
        userName: '',
        fullName: '',
        token: ''
    }
})

export const mutations = {
    increment(state) {
        state.counter++
    },
    add(state, text) {
        state.list.push({
            text,
            done: false
        })
    },
    remove(state, { todo }) {
        state.list.splice(state.list.indexOf(todo), 1)
    },
    toggle(state, todo) {
        todo.done = !todo.done
    },
    login(state, user) {
        console.log('This is the login mutation' + JSON.stringify(user));
        console.log('UserName: ' + user.userName)
        if(user.userName === "alwindoss" && user.password === "opensesame") {
            state.user.isLoggedIn = true;
            state.user.userName = user.userName;
            state.loggedIn = true;
            console.log('valid user');
        } else {
            console.log('Invalid user');
        }
    }
}

export const actions = {
    login(state, user) {
        console.log('This is the login action');
        console.log(JSON.stringify(state));
        console.log(JSON.stringify(user));
        state.commit('login', user)
    },
}