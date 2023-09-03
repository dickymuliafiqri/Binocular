<script setup>
let props = defineProps(["pending", "domains"]);

const columns = [
  {
    key: "domain",
    label: "Domain",
  },
  {
    key: "port",
    label: "Port",
  },
  {
    key: "status_code",
    label: "Status Code",
  },
  {
    key: "webserver",
    label: "Webserver",
  },
  {
    key: "host",
    label: "Host",
  },
  {
    key: "response_time",
    label: "Response Time",
  },
  {
    key: "ping",
    label: "Ping",
  },
];

function pingTest(id, url) {
  for (let i in props.domains) {
    if (props.domains[i].id == id) {
      let startTime = new Date();
      $fetch(url, {
        mode: "no-cors",
        server: false,
      })
        .then(() => {
          props.domains[i].ping = new Date() - startTime;
        })
        .catch(() => {
          props.domains[i].ping = -1;
        });

      break;
    }
  }
}
</script>

<template>
  <UTable :columns="columns" :rows="domains" :loading="pending" :key="domains">
    <template #domain-data="{ row }">
      <a :href="row.raw.url" target="_blank">{{ row.domain }}</a>
    </template>
    <template #status_code-data="{ row }">
      <UBadge
        v-if="row.status_code == 200"
        :ui="{ rounded: 'rounded-full' }"
        :label="row.status_code.toString()"
        color="green"
        variant="subtle"
      />
    </template>
    <template #webserver-data="{ row }">
      <UBadge
        v-if="row.webserver.toLowerCase() == 'cloudflare' || row.webserver.toLowerCase() == 'cloudfront'"
        :ui="{ rounded: 'rounded-full' }"
        :label="row.webserver"
        :color="row.webserver == 'cloudflare' ? 'orange' : 'blue'"
        variant="subtle"
      />
    </template>
    <template #host-data="{ row }">
      <a :href="'http://' + row.host" target="_blank">{{ row.host }}</a>
    </template>
    <template #ping-data="{ row }">
      <UButton
        :icon="row.ping == 0 ? 'i-heroicons-arrow-path' : row.ping < 0 ? 'i-heroicons-x-mark' : null"
        :label="row.ping > 0 ? row.ping + 'ms' : null"
        :color="row.ping == 0 ? 'yellow' : row.ping > 0 ? 'green' : 'red'"
        :ui="{ rounded: 'rounded-full' }"
        size="2xs"
        variant="outline"
        square
        @click="pingTest(row.id, row.raw.url)"
      />
    </template>
  </UTable>
</template>
