<template>
  <div id="editarConsulta">
    <NavbarPaciente/>
    <h1>Remarcar Consulta</h1>
    <form class="form">
       <div class="mb-3">
        <label for="medicoInput" class="form-label">Médico:</label>
        <input type="text" class="form-control" id="medicoInput" v-model="consulta.medicName" disabled/>
      </div>
      <div class="mb-3">
        <label for="dataInput" class="form-label">Data:</label>
        <input type="date" timezone=[[pt-BR]] class="form-control" id="dataInput" v-model="consulta.day" />
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Horário:</label>
        <input type="time" class="form-control" id="horarioInput" v-model="consulta.hourInit" />
      </div>
      <div class="mb-3">
        <label for="precoInput" class="form-label">Preço:</label>
        <input type="text" class="form-control" id="precoInput" v-model="consulta.price" disabled />
      </div>
      <button class="btn-success">Atualizar</button>
      <router-link class="btn-voltar" to="/pacienteHome"><button  class="btn-warning">Voltar</button></router-link>
    </form>
   
  </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'

export default {
  data(){
    return{
      consulta: {
        id: this.$route.params.consulta.id,
        status: this.$route.params.consulta.status,
        effected:  this.$route.params.consulta.effected,
        medicId:  this.$route.params.consulta.medicId,
        clienteId:  this.$route.params.consulta.clienteId,
        medicName: this.$route.params.consulta.medicName,
        clientName: this.$route.params.consulta.clientName,
        state: this.$route.params.consulta.state,
        city: this.$route.params.consulta.city,
        price:  this.$route.params.consulta.price,
        day:  this.formatDate(this.$route.params.consulta.day),
        hourInit: this.formatHour(this.$route.params.consulta.hourInit),
        hourEnd: this.formatHour(this.$route.params.consulta.hourEnd)
      }
    }
  },
  components:{
    NavbarPaciente
  },
  methods:{
    formatDate(d){
      let res = d.split('T')

      return res[0]
    },
     formatHour(hour){
      let res = hour.split('T')
      res = res[1].split(':')
      res = res[0]+':'+res[1]
      return res
    }
  }
}
</script>


<style>
*{
  font-family: Arial, Sans-serif, times;
}

#editarConsulta{
  width: 100vw;
  height: 100vh;
  background-color: var(--primary);
}

h1{
  margin-left: 20px;
  color: #fff;

}

.form{
  width: 50%;
  margin-left: 20px;
  margin-top: 30px;
  padding: 20px 10px;
  border: 2px solid #aaa;
  border-radius: 8px;
  background-color: var(--background-default);
}

.btn-success{
  width: 40%;
  height: 40px;
  border-radius: 6px;
  font-weight: bold;
}

.btn-voltar{
  color: #fff;
}

.btn-warning{
  width: 40%;
  height: 40px;
  border-radius: 6px;
  color: #fff !important;
  font-weight: bold;
  float: right;
}

</style>