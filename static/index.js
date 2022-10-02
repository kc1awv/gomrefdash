import { createApp } from 'https://unpkg.com/vue@3/dist/vue.esm-browser.prod.js'

const refreshTime = 20000;

// station stuff
createApp({
    delimiters: ['${', '}'],
    data() {
        return {
            stationdata: {},
            timer: ''
        }
    },
    methods: {
        fetchStationsList () {
            $.getJSON('/json/stations', (stationdata) => {
                this.stationdata = stationdata
            });
        },
    },
    mounted() {
        this.timer = setInterval(this.fetchStationsList, refreshTime)
        this.fetchStationsList()
    },
}).mount('#stations')

// modules stuff
createApp({
    delimiters: ['${', '}'],
    data() {
        return {
            moduledata: {},
            timer: ''
        }
    },
    methods: {
        fetchModulesList () {
            $.getJSON('/json/modulesinuse', (moduledata) => {
                this.moduledata = moduledata
            });
        },
    },
    mounted() {
        this.timer = setInterval(this.fetchModulesList, refreshTime)
        this.fetchModulesList()
    },
}).mount('#modules')

