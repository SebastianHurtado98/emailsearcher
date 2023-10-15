<template>
  <div class="bg-gray-100 p-4">
    <table class="min-w-full bg-white">
      <thead class="bg-gray-800 text-white">
        <tr>
          <th class="px-2">From</th>
          <th class="px-2">To</th>
          <th class="px-2">Subject</th>
          <th class="px-2">Message</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in results" 
            :key="item.id" 
            @click="onSelect(item)"
            class="hover:bg-gray-300 cursor-pointer">
          <td class="px-2 border-b border-gray-200">{{ truncate(item.from, 10) }}</td>
          <td class="px-2 border-b border-gray-200">{{ truncate(concatenateStrings(item.to), 10)}}</td>
          <td class="px-2 border-b border-gray-200">{{ truncate(item.subject, 20) }}</td>
          <td class="px-2 border-b border-gray-200">{{ truncate(item.message, 100) }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  props: ['results'],
  methods: {
    onSelect(item) {
      this.$emit('select', item);
    },
  truncate(string, length = 20) {
        if (!string || typeof string !== 'string') {
          return '';
        }
        if (string.length > length) {
          return string.slice(0, length) + '...';
        } else {
          return string;
        }
      },
    concatenateStrings(stringArray) {
      if (!Array.isArray(stringArray)) {
        console.error("concatenateStrings method requires an array as parameter");
        return '';
      }
      return stringArray.join(', ');
    }
  }
};
</script>
