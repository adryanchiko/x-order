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
      <v-table>
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
    </div>
    <div>
    </div>
  </div>
</template>

<script lang="ts" setup>

import moment from 'moment'
import { ref, watch, onMounted } from 'vue'

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

const headers = [
  {
    text: 'Dessert (100g serving)',
    align: 'start',
    sortable: false,
    value: 'name',
  },
  { text: 'Calories', value: 'calories' },
  { text: 'Fat (g)', value: 'fat' },
  { text: 'Carbs (g)', value: 'carbs' },
  { text: 'Protein (g)', value: 'protein' },
  { text: 'Iron (%)', value: 'iron' },
]

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
let total_page = ref(0)
let total_records = ref(0)

let keyword = ref('')
let startDate = ref('')
let endDate = ref('')

watch(keyword, () => {
  debounce(doFetch, 500)
})

watch(startDate, () => {
  debounce(doFetch, 500)
})

watch(endDate, () => {
  debounce(doFetch, 500)
})

onMounted(() => {
  doFetch()
})

function newDateFormat(v: string): string {
  const momentDate = moment(v)
  return momentDate.isValid() ? momentDate.format('DD-MMM-YYYY HH:mm:ss') : ''
}

function doFetch() {
  const params = {
    keyword: keyword.value,
    skip: '0',
    start_date: '',
    end_date: ''
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

  fetch('http://localhost:8000/api/v1/orders?' + new URLSearchParams(params).toString())
    .then((response) => {
      return response.json()
    })
    .then(data => {
      records.value = data.records
      total_records.value = data.total_records
      total_page.value = data.total_page
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
