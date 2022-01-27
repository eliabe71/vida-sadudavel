<template>
  <div id="pacienteHistoricoConsultas">
    <NavbarPaciente/>
    <h2 class="title">Meu histórico de Consultas</h2>
    <div class="bloco-consultas">
        <table>
          <tr v-for="(cons,index) in divide(consultas)" :key="index">
            <td v-for="c in cons" :key="c.id">
              <div class="card">
                <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo"> 
                <div class="card-body">
                <h5 class="card-title">Consulta</h5>
                <p class="card-text"><b>Médico:</b> {{c.medicName}} </p>
                <p class="card-text"><b>Local:</b> {{c.city+'-'+c.state}} </p>
                <p class="card-text"><b>Data:</b> {{formatDate(c.day)}}</p>
                <p class="card-text"><b>Horário:</b> {{formatHour(c.hourInit)}}</p>
                <p class="card-text"><b>Preço:</b> R$ {{c.price}},00</p>
                </div>
                <div class="cart-footer">
                  <span class="icon btn-excluir" data-bs-toggle="modal" :data-bs-target="'#confModalexcluir'+c.id">
                      <font-awesome-icon :icon="['fas', 'trash-alt']" />
                  </span>
                </div>
              </div>
              <div class="modal fade" :id="'confModalexcluir'+c.id" tabindex="-1" aria-labelledby="confimationModalLabel" aria-hidden="true">
                  <div class="modal-dialog modal-dialog-centered">
                    <div class="modal-content">
                      <div class="modal-header">
                        <h5 class="modal-title" id="confimationModalLabel">Excluir consulta</h5>
                        <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                      </div>
                      <div class="modal-body">
                        <p>Deseja realmente excluir a consulta do histórico?</p>
                      </div>
                      <div class="modal-footer">
                        <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancelar</button>
                        <button type="button" v-on:click="excluirConsulta(c.id)" data-bs-dismiss="modal" class="btn btn-danger">Excluir {{c.id}}</button>
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

export default {
  data(){
    return {
      consultas: [],
      pacienteLogado: {id:1}
    }
  },
  mounted(){
    Consultas.listarPorPaciente(this.pacienteLogado.id).then(res => {
      this.consultas = res.data.Data.filter(c => {
        return (c.effected === true)
      })
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
      console.log(res)
      return res
    },
    formatDate(d){
      let res = d.split('T')
      res = res[0].split('-')
      res = res[2]+'/'+res[1]+'/'+res[0]
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
          alert("Consulta excluída com sucesso!")
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

#pacienteHistoricoConsultas{
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
  background-color:  var(--primary);
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
  justify-content: center;
}

.icon {
  margin-left: 10px;
  margin-right: 10px;
}

.icon:hover{
  cursor: pointer;
}

.btn-excluir{
  color: red;
}

</style>