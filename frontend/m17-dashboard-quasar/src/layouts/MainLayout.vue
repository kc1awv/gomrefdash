<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated>
      <q-toolbar>
        <q-toolbar-title><a to="/">{{ metadata.reflector_callsign }}</a></q-toolbar-title>
        <q-tabs>
          <q-route-tab to="/" label="Home" />
          <q-route-tab to="/links" label="Links" />
          <q-route-tab to="/peers" label="Peers" />
        </q-tabs>
    <q-btn :label="mode === 'light' ? 'Light Mode' : 'Dark Mode'" color="primary" @click="toggleMode" v-model="$q.dark" />


      </q-toolbar>
    </q-header>

    <q-page-container>
      <router-view />
    </q-page-container>

    <q-footer align="right" elevated class="bg-grey-8 text-white">
      <div class="row">
        <div class="col-12" align="left">
          mrefd Version: {{ metadata.reflector_version }} | Dashboard Version: {{ metadata.dashboard_version }} | Sysop: {{ metadata.sysop_email }} | IP: {{ metadata.ipV4 }} {{ metadata.ipV6 }}
        </div>
      </div>
    </q-footer>
  </q-layout>
</template>

<script>
import { defineComponent, ref } from "vue";
import { useQuasar } from "quasar";

// Import the Axios library
import axios from "axios";
import { store } from "quasar/wrappers";
let metadata_default = {"reflector_callsign": "", "dashboard_version":"","ipV4":"","ipV6":"","":"","reflector_version":"","sysop_email":""};

export default defineComponent({
  name: "MainLayout",

  components: {
  },

  setup() {
    const leftDrawerOpen = ref(false);
    return {
      leftDrawerOpen,
      toggleLeftDrawer() {
        leftDrawerOpen.value = !leftDrawerOpen.value;
      },
    };
  },
  data() {
    this.$q.dark.set("auto");
    let mode = "light";
    if (this.$q.dark.isActive) {
      mode = "dark";
    }
    return {
      mode: mode,
      metadata: metadata_default,
    };
  },
  mounted() {
    this.fetchMetadata();
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
    toggleMode() {
      // Switch between light and dark mode
      this.mode = this.mode === "light" ? "dark" : "light";
      this.$q.dark.toggle();
    },
  },
});
</script>
