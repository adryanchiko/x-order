<template>
  <div class="pa-4">
    <v-text-field v-model="keyword" label="Search" prepend-icon="mdi-magnify" variant="outlined" clearable />
    <p class="font-weight-medium">
      Created Date
    </p>
    <div class="d-block">
      <v-row no-gutters>
        <v-col cols="1">
          <div class="my-2">
            Start
          </div>
        </v-col>
        <v-col>
          <div class="my-2">
            <input v-model="startDate" type="date" id="start" name="trip-start">
          </div>
        </v-col>
      </v-row>
    </div>
    <div class="d-block">
      <v-row no-gutters>
        <v-col cols="1">
          <div class="my-2">
            End
          </div>
        </v-col>
        <v-col>
          <div class="my-2">
            <input v-model="endDate" type="date" id="end" name="trip-end">
          </div>
        </v-col>
      </v-row>
    </div>
    <div class="my-4">
      <p class="font-weight-bold">Total Amount : {{ total_amount > 0 ? '$' + total_amount : '-' }}</p>
    </div>
    <div>
      <v-table class="border">
        <thead>
          <tr>
            <th class="text-left">
              Order Name
            </th>
            <th class="text-left">
              Customer Company
            </th>
            <th class="text-left">
              Customer Name
            </th>
            <th class="text-left">
              Order Date
            </th>
            <th class="text-left">
              Delivered Amount
            </th>
            <th class="text-left">
              Total Amount
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in records">
            <td>
              <p class="font-weight-medium">{{ item.order_name }}</p>
              <p>{{ item.product }}</p>
            </td>
            <td>{{ item.customer_company }}</td>
            <td>{{ item.customer_name }}</td>
            <td>{{ newDateFormat(item.created_at) }}</td>
            <td>{{ item.delivered_amount > 0 ? '$' + item.delivered_amount : '-' }}</td>
            <td>{{ item.total_amount > 0 ? '$' + item.total_amount : '-' }}</td>
          </tr>
        </tbody>
      </v-table>
      <div class="mt-4">
        <v-pagination v-model="page" :length="total_pages" active-color="primary" density="comfortable" />
      </div>
    </div>
    <div>
    </div>
  </div>
</template>

<script lang="ts" setup>

import moment from 'moment'
import { ref, watch, onBeforeMount, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import router from '@/router';

const route = useRoute()

let page = ref(1)
let total_amount = ref(0)
let records = ref([{
  id: '',
  price_per_unit: '',
  quantity: '',
  product: '',
  created_at: '',
  order_name: '',
  customer_name: '',
  customer_company: '',
  delivered_amount: 0,
  total_amount: 0,
}])
let total_pages = ref(0)
let total_records = ref(0)
let keyword = ref('')
let skip = ref('0')
let startDate = ref('')
let endDate = ref('')

watch(keyword, () => {
  debounceFetch()
})

watch(startDate, () => {
  debounceFetch()
})

watch(endDate, () => {
  debounceFetch()
})

watch(page, () => {
  skip.value = ((page.value - 1) * 5).toString()
  doFetch()
})

onBeforeMount(() => {
  loadState()
})

onMounted(() => {
  debounceFetch()
})

function debounce<T>(fn: T, wait: number) {
  const timeoutId = window.setTimeout(() => { }, 0);
  for (let id = timeoutId; id >= 0; id -= 1) {
    window.clearTimeout(id);
  }
  setTimeout(() => {
    if (typeof fn === 'function') {
      fn()
    }
  }, wait)
}

function newDateFormat(v: string): string {
  const momentDate = moment(v)
  return momentDate.isValid() ? momentDate.format('DD-MMM-YYYY HH:mm:ss') : ''
}

function loadState() {
  keyword.value = route.query['keyword']?.toString() || ''
  skip.value = route.query['skip']?.toString() || ''
  startDate.value = route.query['start_date']?.toString() || ''
  endDate.value = route.query['end_date']?.toString() || ''
  page.value = +(route.query['page']?.toString() || 1)
}

function saveState(params: any) {
  router.replace({ path: '/orders', query: params })
}

function debounceFetch() {
  debounce(doFetch, 200)
}

function doFetch() {
  const params = {
    keyword: keyword.value,
    skip: skip.value,
    start_date: '',
    end_date: '',
    page: page.value.toString()
  }

  if (startDate.value) {
    let sDate: Date = new Date(startDate.value)
    sDate.setHours(0, 0, 0)

    params.start_date = sDate.toISOString()
  }

  if (endDate.value) {
    let eDate: Date = new Date(endDate.value)
    eDate.setHours(23, 59, 59)

    params.end_date = eDate.toISOString()
  }

  // save state
  saveState(params)

  fetch('http://localhost:8000/api/v1/orders?' + new URLSearchParams(params).toString())
    .then((response) => {
      return response.json()
    })
    .then(data => {
      records.value = data.records
      total_records.value = data.total_records
      total_pages.value = data.total_pages
    })

  fetch('http://localhost:8000/api/v1/orders-amount?' + new URLSearchParams(params).toString())
    .then((response) => {
      return response.json()
    })
    .then(data => {
      total_amount.value = data
    })
}
</script>

<style>
.border {
  border: 1px solid;
  border-radius: 0.5em;
}
</style>