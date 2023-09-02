<script setup>
defineProps(["domains"]);

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
</script>

<template>
  <UTable :columns="columns" :rows="domains">
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
        v-if="row.webserver == 'cloudflare'"
        :ui="{ rounded: 'rounded-full' }"
        :label="row.webserver"
        color="orange"
        variant="subtle"
      />
      <UBadge
        v-else-if="row.webserver == 'CloudFront'"
        :ui="{ rounded: 'rounded-full' }"
        :label="row.webserver"
        color="blue"
        variant="subtle"
      />
    </template>
    <template #host-data="{ row }">
      <a :href="'http://' + row.host" target="_blank">{{ row.host }}</a>
    </template>
    <template #ping-data="{ row }">
      <UButton
        v-if="row.ping == 0"
        icon="i-heroicons-arrow-path"
        size="2xs"
        color="yellow"
        variant="outline"
        :ui="{ rounded: 'rounded-full' }"
        square
      />
      <UButton
        v-else-if="row.ping > 0"
        :label="row.ping + 'ms'"
        size="2xs"
        color="green"
        variant="outline"
        :ui="{ rounded: 'rounded-full' }"
        square
      />
      <UButton
        v-else
        icon="i-heroicons-x-mark"
        size="2xs"
        color="red"
        variant="outline"
        :ui="{ rounded: 'rounded-full' }"
        square
      />
    </template>
  </UTable>
</template>
