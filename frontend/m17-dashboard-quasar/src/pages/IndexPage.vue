<template>
  <div class="q-pa-sm">
    <div class="row">
      <div class="col-xs-12 col-sm-10 col-md-8">
        <div class="q-pa-md">
          <q-markup-table>
            <thead>
              <tr>
                <th class="gt-xs text-right" style="max-width:50px">#</th>
                <th class="text-center" style="max-width:50px">Flag</th>
                <th class="text-center" style="max-width:150px"><q-btn flat no-caps>Callsign</q-btn></th>
                <th class="gt-xs text-left" style="max-width:50px">Suf</th>
                <th class="gt-xs text-left" style="max-width:100px">Link/Peer</th>
                <th class="text-left" style="max-width:50px">Mod</th>
                <th class="text-left" style="max-width:150px">Last Heard</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="station in station_rows" :key="station">
                <td class="gt-xs text-right" style="max-width:50px">{{ station.id }}</td>
                <td class="text-center" style="max-width:50px"><img width=30 :src="`flags/${station.country.countrycode}.png`"/></td>
                <td class="text-center" style="max-width:150px">
                  <q-btn v-if="station.txactive == true" color="red" text-color="white" :href="`https://www.qrz.com/db/${station.callsign}`" target="_blank"><u>{{ station.callsign }}</u></q-btn>
                  <q-btn icon="volume" v-else flat :href="`https://www.qrz.com/db/${station.callsign}`" target="_blank"><u>{{ station.callsign }}</u></q-btn>
                </td>
                <td class="gt-xs text-left" style="max-width:50px">{{ station.callsignsuffix }}</td>
                <td class="gt-xs text-left" style="max-width:100px">{{ station.vianode }}</td>
                <td class="text-left" style="max-width:50px">{{ station.onmodule }}</td>
                <td class="text-left" style="max-width:150px">{{ station.lastheardlocal }}</td>
              </tr>
            </tbody>
          </q-markup-table>
        </div>
      </div>
      <div class="col-xs-12 col-sm-5 col-md-4">
        <div class="q-pa-sm">
          <q-card>
            <q-card-section>
              <div align="center">Modules In Use</div>
              <div class="q-pa-sm">
                <div class="row">
                  <div
                    class="col-6 q-pa-sm"
                    v-for="module in modules"
                    :key="module.name">
                    <q-card class="mycard">
                      <q-card-section>
                        <q-item-label header align="center"
                          ><div class="text-h6">{{ module.name }}</div>
                        </q-item-label>
                        <q-item-label align="center">
                          <q-list bordered separator>
                            <q-item clickable v-ripple v-for="callsign in module.callsigns" :key="callsign">
                              <q-item-section>{{ callsign }}</q-item-section>
                            </q-item>
                          </q-list>
                        </q-item-label>
                    </q-card-section>
                    </q-card>
                  </div>
                </div>
              </div>
            </q-card-section>
          </q-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";
import { ref } from "vue";
import { localTimeString, isLessThanOneMinuteAgo } from "../js/utilities.js";
import { useQuasar } from "quasar";
import { createArrayExpression } from "@vue/compiler-core";
// Import the Axios library
import axios from 'axios';

export default defineComponent({
  data() {
    return {
      name: "M17 Dashboard",
      station_rows: [],
      modules: [],
      pagination: ref({
        rowsPerPage: 0,
      }),
      // Store the interval ID so we can cancel it later
      intervalId: null,
    };
  },
  mounted() {
    this.fetchMetadata();
    this.fetchStationsList();
    // Fetch and update data every 20 seconds
    this.intervalId = setInterval(() => {
      this.fetchStationsList();
    }, 20000);
  },
  beforeRouteLeave(to, from, next) {
    // Cancel the interval before the component is destroyed
    clearInterval(this.intervalId);

    // Continue to the next hook
    next();
  },
  methods: {
    fetchMetadata() {
      let url = "/json/metadata"
      axios.get(url)
        .then((response) => {
          this.metadata = response.data;
        })
        .catch((error) => {
          // Print any error messages to the console
          console.error(error);
        });
    },
    fetchStationsList() {
      let url = "/json/stations"
      axios.get(url)
        .then((response) => {
          let stationdata = response.data;
          stationdata.stations.forEach(function (station, index) {
            station.id = index + 1;
            station.lastheardlocal = localTimeString(station.lastheard);
            station.txactive = isLessThanOneMinuteAgo(station.lastheard);
            stationdata.stations[index] = station;
          });
          this.station_rows = stationdata.stations;
        })
        .catch((error) => {
          // Print any error messages to the console
          console.error(error);
        });
      let url2 = "/json/modulesinuse"
      axios.get(url2)
        .then((response) => {
          this.modules = response.data;
        })
        .catch((error) => {
          // Print any error messages to the console
          console.error(error);
        });
    },
  },
});
</script>
