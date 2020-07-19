export const state = () => ({
    user: {
        isLoggedIn: false,
        userName: '',
        fullName: '',
        token: ''
    }
})

export const mutations = {
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
        state.user.isLoggedIn = true;
        state.user.userName = user.userName;
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