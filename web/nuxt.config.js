const parseArgs = require("minimist")
const argv = parseArgs(process.argv.slice(2), {
  alias: {
    H: "hostname",
    p: "port"
  },
  string: ["H"],
  unknown: parameter => false
})

const port =
  argv.port ||
  process.env.PORT ||
  process.env.npm_package_config_nuxt_port ||
  "3000"
const host =
  argv.hostname ||
  process.env.HOST ||
  process.env.npm_package_config_nuxt_host ||
  "localhost"

const baseUrl = process.env.BASE_URL || `http://${host}:${port}`
const apiBaseUrl = process.env.API_BASE_URL
const title = 'ホーチミンおすすめカフェ'

module.exports = {
  mode: 'spa',
  env: {
    baseUrl,
    apiBaseUrl,
  },
  head: {
    title: title,
    titleTemplate: `%s - ${title}`,
    meta: [
      { charset: "utf-8" },
      {
        name: "viewport",
        content:
          "width=device-width, initial-scale=1"
      },
      {
        vmid: 'og:title',
        property: 'og:title',
        content: 'Top',
        template: chunk => {
          if (!chunk || chunk === '') {
            return 'ホーチミンおすすめカフェ'
          }
          return `${chunk} - ${title}`
        }
      },
      {
        hid: "description",
        name: "description",
        content: `${title}。ホーチミン在住がリピートするおすすめのカフェを紹介します。`,
      },
      {
        vmid: 'og:description',
        property: 'og:description',
        content: `${title}。ホーチミン在住がリピートするおすすめのカフェを紹介します。`,
        template: '%s'
      },
      {
        vmid: 'og:url',
        property: 'og:url',
        content: baseUrl
      },
      {
        vmid: 'og:type',
        property: 'og:type',
        content: 'website'
      },
      {
        vmid: 'og:locale',
        property: 'og:locale',
        content: 'ja_JP'
      },
      {
        vmid: 'og:image',
        property: 'og:image',
        content: `${baseUrl}/images/ogp.jpg`
      }
    ],
    link: [
      {
        rel: "icon",
        type: "image/x-icon",
        href: "/favicon.ico"
      }
    ]
  },
  plugins: [
    { src: '~/plugins/vue-material' },
    '~/plugins/vue-i18n',
  ],
  /*
  ** Customize the progress-bar color
  */
  loading: { color: "#3B8070" },
  loadingIndicator: {
    name: 'chasing-dots',
    color: '#0288D1',
  },
  render: {
    gzip: {threshold: 9},
    http2: {push: false}
  },
  /*
  ** Build configuration
  */
  css: [
    "~/assets/scss/main.scss",
    'vue-material/dist/vue-material.min.css',
  ],
  build: {},
  modules: [
    "@nuxtjs/axios",
    "~/modules/typescript.js",
    '@nuxtjs/pwa',
  ],
  // axios: {}
}
