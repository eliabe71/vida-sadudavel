<template>
  <div class="container">
    <form class="form">
       <div class="mb-3">
        <label for="medicoInput" class="form-label">Médico:</label>
        <input type="text" class="form-control" id="medicoInput" v-model="consulta.medico.nome" disabled/>
      </div>
      <div class="mb-3">
        <label for="dataInput" class="form-label">Data:</label>
        <input type="date" class="form-control" id="dataInput" v-model="consulta.data" />
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Horário:</label>
        <input type="text" class="form-control" id="horarioInput" v-model="consulta.horario" />
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Preço:</label>
        <input type="text" class="form-control" id="horarioInput" v-model="consulta.preco" disabled />
      </div>
    </form>
  </div>
</template>

<script>
  export default {
    data(){
      return{
        consulta: this.$route.params.consulta
      }
    }
  }
</script>


<style >

</style>