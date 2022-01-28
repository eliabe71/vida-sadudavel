<template>
  <div id="pacienteSolicitacoes">
    <NavbarPaciente/>
    <h2 class="title">Minhas Solicitações de Consultas</h2>
    <div class="bloco-consultas">
      <table>
        <tr v-for="(cons,index) in divide(consultas)" :key="index">
          <td v-for="c in cons" :key="c.id">
            <div class="card">
              <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo"> 
              <div class="card-body">
                <h5 class="card-title">Consulta</h5>
                <p class="card-text"><b>Médico:</b> {{c.medicName}}</p>
                <p class="card-text"><b>Local:</b> {{c.city+'-'+c.state}} </p>
                <p class="card-text"><b>Data:</b> {{formatDate(c.day)}}</p>
                <p class="card-text"><b>Horário:</b> {{formatHour(c.hourInit)}}</p>
                <p class="card-text"><b>Preço:</b> R$ {{c.price}},00</p>
              </div>
              <div class="cart-footer">
                <button class="btn-danger btn-excluir" data-bs-toggle="modal" :data-bs-target="'#confModalexcluir'+c.id">Cancelar solicitação</button>
              </div>
              <div class="modal fade" :id="'confModalexcluir'+c.id" tabindex="-1" aria-labelledby="confimationModalLabel" aria-hidden="true">
                <div class="modal-dialog modal-dialog-centered">
                  <div class="modal-content">
                    <div class="modal-header">
                      <h5 class="modal-title" id="confimationModalLabel">Excluir consulta</h5>
                      <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                    </div>
                    <div class="modal-body">
                      <p>Deseja realmente excluir a solicitação de consulta?</p>
                    </div>
                    <div class="modal-footer">
                      <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancelar</button>
                      <button type="button" v-on:click="excluirConsulta(c.id)" data-bs-dismiss="modal" class="btn btn-danger">Excluir</button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Consultas from '../../services/consultas'
import Pacientes from '../../services/pacientes'

export default {
  data(){
    return {
      consultas: [],
      pacienteLogado: {},
    }
  },
  mounted(){
    Consultas.listarPorPaciente(1).then(res => {
      this.consultas = res.data.Data.filter(c => {
        return (c.status === false && c.effected === false)
      })
    }),
    Pacientes.getPaciente(1).then(res=>{
      this.pacienteLogado = res.data.Data[0]
    })
  },
  methods: {
    divide(cons){
      let res = []
      let aux = []
      for(let i=0; i<cons.length; i++){
        aux.push(cons[i])
        if(((i+1)%4 === 0 && i !== 0) || i === cons.length-1){
          res.push(aux)
          aux = []
        }
      }
      return res
    },
  
    formatDate(d){
      let res = d.split('T')
      res = res[0].split('-')
      res = res[2]+'/'+res[1]+'/'+res[0]
      return res
    },
    formatDateJson(d){
      let res = d.split('T')
      res = res[0].split('-')
      res = res[0]+'-'+res[1]+'-'+res[2]
      return res
    },
    formatHour(hour){
      let res = hour.split('T')
      res = res[1].split(':')
      res = res[0]+':'+res[1]
      return res
    },
    excluirConsulta(id){
      Consultas.excluir(id).then(res => {
        if(res.status === 200){
          this.$router.go()
          alert("Solicitção excluída com sucesso!")
        }else
          alert("Erro ao excluir consulta!")

      })
    },
  },

  components:{
    NavbarPaciente
  }
}
</script>

<style>

*{
  font-family: Arial, Sans-serif, times;
}

#pacienteSolicitacoes {
  width: 100vw;
  height: 100vh;
  background-color:  var(--primary);
}

.title{
  color: #fff;
  margin-top: 10px;
  margin-left: 20px;
  margin-bottom: 0;
  font-size: 25pt;
  font-weight: bold;
}

.bloco-consultas{
  display: flex;
}

.bloco-consultas table{
  margin-left: 10px;
  margin-right: 10px;
}

.bloco-consultas td{
  min-width: 18vw;
  padding: 10px;
}

.card{
  background-color: var(--background-default);
  display: flex;
  width: 100%;
  height: 100%;
}

.card img{
  width: 100%;
  height: 100%;
  margin-bottom: 0vw;
}

.card-body{
  margin-top: 0;
  width: 100%;
  height: 100%;
}

.card-title{
  font-size: 2vw;
  font-weight: bold;
  margin-bottom: 0.1vw;
  margin-top: 0;
}
.card-text{
  font-size: 1.5vw;
  margin-bottom: 0.05vw;
  margin-top: 0;
}

.card-footer{
  width: 100%;
  height: 3vw;
  display: flex;
  align-items: center;
}

.btn-excluir{
  margin-left: 0.8vw;
  margin-bottom: 10px;
  border-radius: 5px;
}

</style>