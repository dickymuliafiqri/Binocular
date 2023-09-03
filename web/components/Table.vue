<script setup>
let props = defineProps(["pending", "domains"]);

const page = ref(1);
const listPerPage = 10;

const data = computed(() => {
  return props.domains.slice((page.value - 1) * listPerPage, page.value * listPerPage);
});

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

function pingTest(url) {
  for (let i in props.domains) {
    if (props.domains[i].raw.url == url) {
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
  <UTable :columns="columns" :rows="data" :loading="pending" :key="data">
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
      <UTooltip text="Ping Test">
        <UButton
          :icon="row.ping == 0 ? 'i-heroicons-arrow-path' : row.ping < 0 ? 'i-heroicons-x-mark' : null"
          :label="row.ping > 0 ? row.ping + 'ms' : null"
          :color="row.ping == 0 ? 'yellow' : row.ping > 0 ? 'green' : 'red'"
          :ui="{ rounded: 'rounded-full' }"
          size="2xs"
          variant="outline"
          square
          @click="pingTest(row.raw.url)"
        />
      </UTooltip>
    </template>
  </UTable>
  <div class="flex w-full justify-end my-3 px-5">
    <UPagination
      v-model="page"
      :page-count="listPerPage"
      :total="domains.length"
      :ui="{ rounded: 'first-of-type:rounded-s-md last-of-type:rounded-e-md' }"
      color="black"
      size="xs"
    >
      <template #prev="{ onClick }">
        <UTooltip text="Previous page">
          <UButton
            icon="i-heroicons-arrow-small-left-20-solid"
            color="primary"
            :ui="{ rounded: 'rounded-full' }"
            class="rtl:[&_span:first-child]:rotate-180 me-2"
            @click="onClick"
            size="2xs"
          />
        </UTooltip>
      </template>
      <template #next="{ onClick }">
        <UTooltip text="Next page">
          <UButton
            icon="i-heroicons-arrow-small-right-20-solid"
            color="primary"
            :ui="{ rounded: 'rounded-full' }"
            class="rtl:[&_span:last-child]:rotate-180 ms-2"
            @click="onClick"
            size="2xs"
          />
        </UTooltip> </template
    ></UPagination>
  </div>
</template>
