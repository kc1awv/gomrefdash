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
                    :key="module.name"
                  >
                    <q-card class="mycard">
                      <q-card-section>
                        <q-item-label header align="center"
                          ><div class="text-h6">{{ module.name }}</div>
                        </q-item-label>
                        <q-item-label
                          v-for="callsign in module.callsigns"
                          :key="callsign"
                        >
                          <q-item-section>{{ callsign }}</q-item-section>
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

const mockStationData = {
  stations: [
    {
      callsign: "AC8ZD",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T16:45:29Z",
    },
    {
      callsign: "W5TMK",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T13:16:07Z",
    },
    {
      callsign: "KI8M",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T13:13:24Z",
    },
    {
      callsign: "KE8VJT",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T12:23:05Z",
    },
    {
      callsign: "W8EAP",
      callsignsuffix: "D",
      vianode: "W8EAP   D",
      onmodule: "B",
      lastheard: "2022-11-26T04:00:36Z",
    },
    {
      callsign: "W8RJD",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T02:40:30Z",
    },
    {
      callsign: "W7BAZ",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-26T00:39:02Z",
    },
    {
      callsign: "AC8ZD",
      callsignsuffix: "D",
      vianode: "AC8ZD D",
      onmodule: "B",
      lastheard: "2022-11-25T23:11:55Z",
    },
    {
      callsign: "AC8ZD",
      callsignsuffix: " ",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T21:24:23Z",
    },
    {
      callsign: "KB8GHQ",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T21:24:08Z",
    },
    {
      callsign: "W8SOX",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T20:50:25Z",
    },
    {
      callsign: "W1IK",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T16:04:15Z",
    },
    {
      callsign: "W8TVO",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T16:03:48Z",
    },
    {
      callsign: "VE3TCS",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T15:16:03Z",
    },
    {
      callsign: "M7CIZ",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T15:08:50Z",
    },
    {
      callsign: "N8ZA",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T12:56:10Z",
    },
    {
      callsign: "KO4PDI",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T00:37:26Z",
    },
    {
      callsign: "KT7TT",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T00:33:58Z",
    },
    {
      callsign: "N8HUB",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T00:31:56Z",
    },
    {
      callsign: "VK4MSS",
      callsignsuffix: "D",
      vianode: "AC8ZD   D",
      onmodule: "A",
      lastheard: "2022-11-25T00:29:17Z",
    },
  ],
};

const MockModulesInUse = [
  { name: "A", callsigns: ["AC8ZD   D"] },
  { name: "B", callsigns: ["AC8ZD D"] },
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
    };
  },
  created() {
    this.$watch(
      () => this.$route.params.name,
      (toName, previousName) => {
        // react to route changes...
        this.name = toName;
        this.tableTitle = toName;
        this.station_columns = this.station_columns();

        //this.station_rows = [];
        this.fetchStationsList();
      }
    );
    this.fetchStationsList();
    this.modules = MockModulesInUse;
    const $q = useQuasar();
    $q.dark.set(false);
  },
  methods: {
    /*fetchStationsList() {
      $.getJSON("{{ .config.SubPath }}/json/stations", (stationdata) => {
        if (
          !stationdata.hasOwnProperty("stations") ||
          stationdata.stations == null
        ) {
          // no stations, so do nothing
          return;
        }

        stationdata.stations.forEach(function (station, index) {
          station.lastheardlocal = localTimeString(station.lastheard);
          stationdata.stations[index] = station;
        });
        this.stationdata = stationdata;
      });
    },*/
    fetchStationsList() {
      let stationdata = mockStationData;
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
    },
  },
});
</script>
