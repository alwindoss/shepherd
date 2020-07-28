import colors from 'vuetify/es5/util/colors'

export default {
  /*
  ** Nuxt rendering mode
  ** See https://nuxtjs.org/api/configuration-mode
  */
  mode: 'universal',
  /*
  ** Nuxt target
  ** See https://nuxtjs.org/api/configuration-target
  */
  target: 'static',
  /*
  ** Headers of the page
  ** See https://nuxtjs.org/api/configuration-head
  */
  head: {
    titleTemplate: '%s - ' + process.env.npm_package_name,
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Global CSS
  */
  css: [
  ],
  /*
  ** Plugins to load before mounting the App
  ** https://nuxtjs.org/guide/plugins
  */
  plugins: [
  ],
  /*
  ** Auto import components
  ** See https://nuxtjs.org/api/configuration-components
  */
  components: true,
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    // '@nuxtjs/eslint-module',
    '@nuxtjs/vuetify'
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://axios.nuxtjs.org/usage
    '@nuxtjs/axios',
    '@nuxtjs/auth',
    '@nuxtjs/pwa',
    // Doc: https://github.com/nuxt/content
    '@nuxt/content'
  ],
  /*
  ** Auth related configurations
  */
  auth: {
    strategies: {
      local: {
        endpoints: {
          login: { url: '/auth/login', method: 'post', propertyName: 'token' },
          logout: { url: '/auth/login', method: 'delete' },
          user: { url: '/auth/user', method: 'get', propertyName: 'user' }
        },
        // tokenRequired: true,
        // tokenType: 'bearer',
        // globalToken: true,
        // autoFetchUser: true
      },
      github: {
        client_id: '0cc8af08088441c7eb55',
        client_secret: '3fdcab531b99b19be140144c772bdc03b5faf36d'
      },
      // social: {
      //   _scheme: 'oauth2',
      //   authorization_endpoint: 'https://accounts.google.com/o/oauth2/auth',
      //   userinfo_endpoint: 'https://www.googleapis.com/oauth2/v3/userinfo',
      //   scope: ['openid', 'profile', 'email'],
      //   access_type: undefined,
      //   access_token_endpoint: undefined,
      //   response_type: 'token',
      //   token_type: 'Bearer',
      //   redirect_uri: undefined,
      //   client_id: '956094645609-ckjhmjvl19c5lvkbt6bcb842lod7pb0i.apps.googleusercontent.com',
      //   token_key: 't7IMG6orchX9vS-QRB_uQg_J',
      //   state: 'UNIQUE_AND_NON_GUESSABLE'
      // },
      // google: {
      //   client_id: '956094645609-ckjhmjvl19c5lvkbt6bcb842lod7pb0i.apps.googleusercontent.com'

      // }
    }
  },
  /*
  ** Axios module configuration
  ** See https://axios.nuxtjs.org/options
  */
  axios: {
    baseURL: 'http://192.168.55.15:8080/api'
  },
  /*
  ** Content module configuration
  ** See https://content.nuxtjs.org/configuration
  */
  content: {},
  /*
  ** vuetify module configuration
  ** https://github.com/nuxt-community/vuetify-module
  */
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    theme: {
      dark: true,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },
  /*
  ** Build configuration
  ** See https://nuxtjs.org/api/configuration-build/
  */
  build: {
  },
  server: {
    port: 3000,
    host: '0.0.0.0'
  },

  router: {
    middleware: ['auth'],
    base: '/shepherd'
  }
}
