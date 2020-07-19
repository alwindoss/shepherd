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
    }
}

export const actions = {
    login(state, user) {
        console.log('This is the login action');
        console.log(JSON.stringify(state));
        console.log(JSON.stringify(user));
    },
}