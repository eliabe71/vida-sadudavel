<template>
  <div id="editarConsulta">
    <form class="form">
       <div class="mb-3">
        <label for="medicoInput" class="form-label">Médico:</label>
        <input type="text" class="form-control" id="medicoInput" v-model="consulta.medico.nome" disabled/>
      </div>
      <div class="mb-3">
        <label for="dataInput" class="form-label">Data:</label>
        <input type="date" timezone=[[pt-BR]] class="form-control" id="dataInput" v-model="consulta.data" />
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Horário:</label>
        <input type="time" class="form-control" id="horarioInput" v-model="consulta.horario" />
      </div>
      <div class="mb-3">
        <label for="precoInput" class="form-label">Preço:</label>
        <input type="text" class="form-control" id="precoInput" v-model="consulta.preco" disabled />
      </div>
      <button class="btn-success">Atualizar</button>
      <button class="btn-warning">Voltar</button>
    </form>
   
  </div>
</template>

<script>
  export default {
    data(){
      return{
        consulta: this.$route.params.consulta 
      }
    },
  }
</script>


<style>

.form{
  width: 50%;
  margin-left: 20px;
  margin-top: 20px;
  padding: 20px 10px;
  border: 2px solid #aaa;
  border-radius: 8px;
}

.btn-success{
  width: 40%;
  height: 40px;
  border-radius: 6px;
}

.btn-warning{
  width: 40%;
  height: 40px;
  border-radius: 6px;
  float: right;
}

</style>