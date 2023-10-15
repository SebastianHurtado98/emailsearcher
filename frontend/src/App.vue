<template>
  <div id="app" class="bg-gray-100 min-h-screen">
    <div class="container mx-auto p-4">
      <SearchBar @search="handleSearch" />
      <div class="content flex mt-4">
        <ResultsTable :results="results" @select="handleSelect" class="flex-1" />
        <ResultDetail :detail="selectedResult" class="flex-1" />
      </div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';
import SearchBar from './components/SearchBar.vue';
import ResultsTable from './components/ResultsTable.vue';
import ResultDetail from './components/ResultDetail.vue';

export default {
  components: {
    SearchBar,
    ResultsTable,
    ResultDetail
  },
  data() {
    return {
      results: [],
      query: '',
      selectedResult: null
    };
  },
  methods: {
    async handleSearch(query) {
      try {
        this.query = query;
        const response = await axios.get('http://localhost:8000/emails/', {
          params: {
            message: query
          }
        });
        this.results = response.data;
        console.log(response.data)
      } catch (error) {
        console.error('An error occurred while fetching data:', error);
      }
    },
    handleSelect(item) {
      this.selectedResult = item;
    }
  }
};
</script>

<style>
</style>
