<template>
  <div id="pacienteAgendarConsulta" @submit="abrirModal()">
    <NavbarPaciente/>
    <h3>Agendar Consulta</h3>
    <form class="form" >
      <div class="mb-3">
        <label for="medicoInput" class="form-label">Médico:</label>
        <input type="text" timezone=[[pt-BR]] class="form-control" v-model="consulta.medicName" id="medicoInput" disabled/>
      </div>
      <div class="mb-3">
        <label for="dataInput" class="form-label">Data:</label>
        <input type="date" timezone=[[pt-BR]]  class="form-control" id="dataInput" v-model="consulta.day" required />
      </div>
      <div class="mb-3">
        <label for="horarioInput" class="form-label">Horário:</label>
        <input type="time" class="form-control" id="horarioInput" v-model="consulta.hourInit" required/>
      </div>
      <div class="mb-3">
        <label for="precoInput" class="form-label">Preço:</label>
        <input type="text" class="form-control" id="precoInput" v-model="consulta.price" disabled />
      </div>
      <button type="submit" class="btn-success" >Agendar</button>
      <router-link class="btn-voltar" to="/paciente/verMedicos"><button  class="btn-warning">Voltar</button></router-link>
    </form>
  </div>
    
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Pacientes from '../../services/pacientes'
import Consultas from '../../services/consultas'

export default {
  data(){
    return{
      consulta: {
        status: false,
        effected:  false,
        medicId:  null,
        clientId: null,
        medicName: null,
        price:  null,
        day:  null,
        hourInit: null,
        hourEnd: null,
      },

      pacienteLogado: {}
    }
  },
  components:{
      NavbarPaciente
  },
  mounted(){
    Pacientes.getPaciente(1).then(res => {
      this.pacienteLogado = res.data.Data[0]
      this.consulta.clientId = res.data.Data[0].id
    }),
    this.init()
    var today = new Date();
    var dd = today.getDate();
    var mm = today.getMonth() + 1; //January is 0!
    var yyyy = today.getFullYear();

    if (dd < 10) {
      dd = '0' + dd;
    }

    if (mm < 10) {
      mm = '0' + mm;
    } 
        
    today = yyyy + '-' + mm + '-' + dd;
    document.getElementById("dataInput").setAttribute("min", today);
  },
  methods: {
    abrirModal(){
      let res = confirm("Deseja realmente marcar essa consulta?")
      if(res === true){
        this.agendarConsulta()
      }
    },
    init () {
      this.consulta.medicId = this.$route.params.medicId
      this.consulta.medicName = this.$route.params.medicName
      this.consulta.price = this.$route.params.price
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
    agendarConsulta(){
      let cons = this.consulta

      cons.hourEnd = this.add25(cons.hourInit)

      Consultas.cadastrar(cons).then(res => {
        if(res.status === 200){
          alert("Consulta marcada com sucesso!")
          this.$router.push('/pacienteHome')
        }
        else
          alert("Erro ao marcar consulta!")
      })
    }
  }
}
</script>


<style>

*{
  font-family: Arial, Sans-serif, times;
}

#pacienteAgendarConsulta {
  width: 100vw;
  height: 100vh;
  background-color:  var(--primary);
}

h3{
  margin-left: 20px;
  color: #fff;
  margin-top: 10px;
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