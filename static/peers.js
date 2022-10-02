import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.prod.js'

const refreshTime = 20000;

// peers stuff
createApp({
    delimiters: ['${', '}'],
    data() {
        return {
            peersdata: {},
            timer: ''
        }
    },
    methods: {
        fetchPeersList () {
            $.getJSON('/json/peers', (peersdata) => {
                this.peersdata = peersdata
            });
        },
    },
    mounted() {
        this.timer = setInterval(this.fetchPeersList, refreshTime)
        this.fetchPeersList()
    },
}).mount('#peers')