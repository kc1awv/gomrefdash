<template>
  <div class="q-pa-sm">
    <div class="row">
      <div class="col-xs-12 col-md-8">
        <div class="q-pa-md">
          <q-table
            :rows="station_rows"
            :columns="station_columns"
            row-key="name"
            virtual-scroll
            v-model:pagination="pagination"
            :rows-per-page-options="[0]"
          />
        </div>
      </div>
      <div class="col-xs-12 col-md-4">
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
import { localTimeString } from "../js/utilities.js";
import { useQuasar } from "quasar";
import { createArrayExpression } from "@vue/compiler-core";
// Import the Axios library
import axios from 'axios';


const station_columns = [
  {
    name: "id",
    align: "right",
    label: "#",
    field: "id",
    sortable: true,
  },
  {
    name: "callsign",
    label: "Callsign",
    field: "callsign",
    sortable: true,
  },
  {
    name: "callsignsuffix",
    label: "Suffix",
    field: "callsignsuffix",
    sortable: false,
  },
  {
    name: "vianode",
    label: "Link/Peer",
    field: "vianode",
    sortable: true,
  },
  {
    name: "onmodule",
    label: "Module",
    field: "onmodule",
    sortable: true,
  },
  {
    name: "lastheardlocal",
    label: "Last Heard",
    field: "lastheardlocal",
    sortable: true,
  },
];

export default defineComponent({
  data() {
    return {
      name: "M17 Dashboard",
      station_columns: station_columns,
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
