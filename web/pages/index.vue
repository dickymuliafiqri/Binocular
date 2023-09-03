<script setup>
let pending = ref(false);
let domains = ref([]);
const inputValue = ref("");

async function getDomainsData() {
  const domain = inputValue.value;
  if (domain != "") {
    pending.value = true;
    domains.value = await $fetch(`http://127.0.0.1:8080/subfinder?domain=${domain}`)
      .then((data) => {
        data = JSON.parse(data);
        for (let i in data.result) {
          data.result[i].ping = 0;
        }
        return data.result.filter((res) => {
          return res.status_code > 0;
        });
      })
      .finally(() => {
        pending.value = false;
      });
  }
}
</script>

<template>
  <div class="flex flex-col w-full pt-6">
    <div class="flex justify-center w-full text-gray-300">
      <UInput v-model="inputValue" placeholder="Search..." variant="none" />
      <UButton icon="i-heroicons-magnifying-glass-solid" :ui="{ rounded: 'rounded-full' }" @click="getDomainsData" />
    </div>
    <div class="mt-5">
      <div class="border rounded border-gray-800 overflow-x-scroll">
        <Table :domains="domains" :pending="pending" :key="pending" />
      </div>
    </div>
  </div>
</template>
