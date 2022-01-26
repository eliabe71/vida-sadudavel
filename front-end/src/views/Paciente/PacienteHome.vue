<template>
    <div id="pacienteHome">
      <NavbarPaciente/>
      <h2 class="title">Minhas Consultas</h2>
      <div class="bloco-consultas">
        <table>
          <tr v-for="(cons,index) in divide(consultas)" :key="index">
            <td v-for="c in cons" :key="c.id">
              <div class="card">
                <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo"> 
                <div class="card-body">
                  <h4 class="card-title">Consulta</h4>
                  <p class="card-text"><b>Médico:</b>
                  {{ searchMedic(c.medicId)}} 
                  </p>
                  <p class="card-text"><b>Local:</b> </p>
                  <p class="card-text"><b>Data:</b> {{formatDate(c.day)}}</p>
                  <p class="card-text"><b>Horário:</b> {{formatHour(c.hourInit)}}</p>
                  <p class="card-text"><b>Preço:</b> R$ {{c.price}},00</p>
                </div>
                <div class="cart-footer">
                  <span class="icon btn-excluir" data-bs-toggle="modal" data-bs-target="#confimationModal">
                    <font-awesome-icon :icon="['fas', 'trash-alt']" />
                  </span>
                  <router-link class="icon btn-editar" :to="{name: 'pacienteEditarConsulta',  params:{id:c.id, consulta:c} }">
                    <font-awesome-icon :icon="['fas', 'edit']" />
                  </router-link>
                </div>
              </div>
            </td>
          </tr>
        </table>
    </div>
      <div class="modal fade" id="confimationModal" tabindex="-1" aria-labelledby="confimationModalLabel" aria-hidden="true">
        <div class="modal-dialog modal-dialog-centered">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title" id="confimationModalLabel">Excluir consulta</h5>
              <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <p>Deseja realmente excluir a consulta?</p>
            </div>
            <div class="modal-footer">
              <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancelar</button>
              <button type="button" class="btn btn-danger">Excluir</button>
            </div>
          </div>
        </div>
      </div>
    </div>
</template>

<script>
import NavbarPaciente from '../../components/NavbarPaciente/NavbarPaciente.vue'
import Consultas from '../../services/consultas'
import Medicos from '../../services/medicos'

export default {
  data(){
    return {
      consultas: [],
      pacienteLogado: {id: 1},
    }
  },
  mounted(){
    Consultas.listarPorPaciente(this.pacienteLogado.id).then(res => {
      this.consultas = res.data.Data.filter(c => {
        return (c.status === true && c.effected === false)
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
      return res
    },
    async searchMedic(id){

      return await Medicos.getMedico(id).then(r => {
        console.log(r.data.Data[0])
        return r.data.Data[0]
      })
        
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
    }
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

#pacienteHome {
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
  justify-content: center;
}

.icon {
  margin-left: 10px;
  margin-right: 10px;
  font-size: 20pt;
}

.icon:hover{
  cursor: pointer;
}

.btn-editar{
  color: var(--primary);
}

.btn-excluir{
  color: red;
}
</style>