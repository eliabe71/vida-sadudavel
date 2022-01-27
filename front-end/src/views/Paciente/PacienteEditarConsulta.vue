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
        <input type="date" timezone=[[pt-BR]] class="form-control" id="dataInput" v-model="consulta.day" requiered/>
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Horário:</label>
        <input type="time" class="form-control" id="horarioInput" v-model="consulta.hourInit" required/>
      </div>
      <div class="mb-3">
        <label for="precoInput" class="form-label">Preço:</label>
        <input type="text" class="form-control" id="precoInput" v-model="consulta.price" disabled />
      </div>
      <button type="submit" class="btn-success" v-on:submit="atualizarConsulta()">Atualizar</button>
      <router-link class="btn-voltar" to="/pacienteHome"><button  class="btn-warning">Voltar</button></router-link>
    </form>
  </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Consultas from '../../services/consultas'

export default {
  data(){
    return{
      consulta: {
        id: null,
        status: null,
        effected:  null,
        medicId:  null,
        clientId:  null,
        medicName: null,
        clientName: null,
        state: null,
        city: null,
        price:  null,
        day:  null,
        hourInit: null,
        hourEnd: null
      }
    }
  },
  components:{
    NavbarPaciente
  },
  mounted(){
      this.consulta.id = this.$route.params.consulta.id
      this.consulta.status = this.$route.params.consulta.status
      this.consulta.effected =  this.$route.params.consulta.effected
      this.consulta.medicId = this.$route.params.consulta.medicId
      this.consulta.clientId = this.$route.params.consulta.clientId
      this.consulta.medicName = this.$route.params.consulta.medicName
      this.consulta.clientName = this.$route.params.consulta.clientName
      this.consulta.state = this.$route.params.consulta.state
      this.consulta.city = this.$route.params.consulta.city
      this.consulta.price =  this.$route.params.consulta.price
      this.consulta.day =  this.formatDate(this.$route.params.consulta.day)
      this.consulta.hourInit = this.formatHour(this.$route.params.consulta.hourInit)
      this.consulta.hourEnd = this.formatHour(this.$route.params.consulta.hourEnd)
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
    },
    add25(h){
      let aux = h.split(':')
      if(aux[1] < 35){
        let b = parseInt(aux[1])+25
        if(b < 10)
          return aux[0]+':0'+(b)
        else
          return aux[0]+':'+(b)
      }
      else{
        if(aux[0]<23){
          let a = parseInt(aux[0])+1
          let b = (parseInt(aux[1])+25)%60
          if(a < 10){
            if(b<10)
              return '0'+a+':0'+(b)
            else
              return '0'+a+':'+(b)
          }else{
            if(b<10)
              return a+':0'+(b)
            else
              return a+':'+(b)
          }
        }
        else{
          let b = (parseInt(aux[1])+25)%60
          if(b<10)
            return '00'+':0'+b
          else
            return '00'+':'+b
        }
      }
    },
    atualizarConsulta(){
      let cons = this.consulta
      cons.status = false

      cons.hourEnd = this.add25(cons.hourInit)

      Consultas.atualizar(cons).then(res => {
        if(res.status === 200){
          alert("Consulta remarcada com sucesso!")
          this.$router.push('/pacienteHome')
        }
        else
          alert("Erro ao remarcar consulta!")
      })
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