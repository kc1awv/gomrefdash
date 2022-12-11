<template>
  <div class="q-pa-sm">
    <div class="row">
      <div class="col-xs-12 col-md-12">
        <div class="q-pa-md">
          <q-table
            :rows="peer_rows"
            :columns="peer_columns"
            row-key="name"
            virtual-scroll
            v-model:pagination="pagination"
            :rows-per-page-options="[0]"
          />
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


const peer_columns = [
  {
    name: "callsign",
    label: "Callsign",
    field: "callsign",
    sortable: true,
  },
  {
    name: "ip",
    label: "IP",
    field: "ip",
    sortable: true,
  },
  {
    name: "linkedmodule",
    label: "Linked Module",
    field: "linkedmodule",
    sortable: true,
  },
  {
    name: "connecttime",
    label: "Connect Time",
    field: "connecttime",
    sortable: true,
  },
  {
    name: "lastheardtime",
    label: "Last Heard Time",
    field: "lastheardtime",
    sortable: true,
  },
];

export default defineComponent({
  data() {
    return {
      name: "Peers",
      peer_columns: peer_columns,
      peer_rows: [],
      pagination: ref({
        rowsPerPage: 0,
      }),
      intervalId: null,
    };
  },
  mounted() {
    this.fetchPeers();
    this.intervalId = setInterval(() => {
      this.fetchPeers();
    }, 20000);
  },
  beforeRouteLeave(to, from, next) {
    // Cancel the interval before the component is destroyed
    clearInterval(this.intervalId);

    // Continue to the next hook
    next();
  },
  methods: {
    fetchPeers() {
      let url = "/json/peers"
      axios.get(url)
        .then((response) => {
          let peers = response.data;
          peers.forEach(function (peer, index) {
            peer.id = index + 1;
            peer.connecttime = localTimeString(peer.connecttime);
            peer.lastheardtime = localTimeString(peer.lastheardtime);
            peers[index] = peer;
          });
          this.peer_rows = peers;
        })
        .catch((error) => {
          // Print any error messages to the console
          console.error(error);
        });
    },
  },
});
</script>
