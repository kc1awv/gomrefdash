import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.prod.js'

const refreshTime = 20000;

// links stuff
createApp({
    delimiters: ['${', '}'],
    data() {
        return {
            linksdata: [],
            timer: ''
        }
    },
    methods: {
        fetchLinksList () {
            $.getJSON('/json/links', (linksdata) => {
                this.linksdata = linksdata
            });
        },
    },
    mounted() {
        this.timer = setInterval(this.fetchLinksList, refreshTime)
        this.fetchLinksList()
    },
}).mount('#links')
