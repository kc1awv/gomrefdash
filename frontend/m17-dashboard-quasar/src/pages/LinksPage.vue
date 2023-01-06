<template>
  <div class="q-pa-sm">
    <div class="row">
      <div class="col-xs-12 col-md-12">
        <div class="q-pa-md">
          <q-table
            :rows="link_rows"
            :columns="link_columns"
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


const link_columns = [
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
    name: "protocol",
    label: "Protocol",
    field: "protocol",
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
      name: "Links",
      link_columns: link_columns,
      link_rows: [],
      pagination: ref({
        rowsPerPage: 0,
      }),
      // Store the interval ID so we can cancel it later
      intervalId: null,
    };
  },
  mounted() {
    this.fetchLinks();
    this.intervalId = setInterval(() => {
      this.fetchLinks();
    }, 20000);
  },
  beforeRouteLeave(to, from, next) {
    // Cancel the interval before the component is destroyed
    clearInterval(this.intervalId);

    // Continue to the next hook
    next();
  },
  methods: {
    fetchLinks() {
      let url = "/json/links"
      axios.get(url)
        .then((response) => {
          let links = response.data;
          links.forEach(function (link, index) {
            link.id = index + 1;
            link.connecttime = localTimeString(link.connecttime);
            link.lastheardtime = localTimeString(link.lastheardtime);
            links[index] = link;
          });
          this.link_rows = links;
        })
        .catch((error) => {
          // Print any error messages to the console
          console.error(error);
        });
    },
  },
});
</script>
