<template>
  <div :class="{'bg-gray-100 p-4': detail, 'p-4': true}">
    <div v-if="detail">
      <h2 class="text-xl font-bold text-gray-700">{{ detail.subject }}</h2>
      <h2 class="text-l font-bold text-gray-700">From: {{ detail.from }}</h2>
      <h2 class="text-l font-bold text-gray-700">To: {{ detail.to }}</h2>
      <p class="mt-2 text-gray-700">{{ detail.message }}</p>
    </div>
    <div v-else>
      <p class="text-gray-500">Seleccione un elemento para ver los detalles</p>
    </div>
  </div>
</template>

<script>
export default {
  props: ['detail']
};
</script>

<style scoped>
/* Puedes añadir estilos específicos del componente aquí si lo necesitas */
</style>
