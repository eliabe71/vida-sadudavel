<template>
  <div id="medicoSolicitacoes">
    <NavbarMedico/>
    <div class="bloco-consultas">
      <table>
        <tr v-for="(solis,index) in divide(solicitacoes)" :key="index">
          <td v-for="s in solis" :key="s.id">
            <div class="card">
              <img src="../../assets/images/logo.png" class="card-img-top" alt="Logo">
              <div class="card-body">
                <h4 class="card-title">Solicitação de consulta</h4>
                <p class="card-text">Paciente: </p>
                <p class="card-text">Data: {{s.day}}</p>
                <p class="card-text">Horário: </p>
              </div>
              <div class="cart-footer">
                <button class="btn-success">Aceitar</button>
              </div>
            </div>
          </td>
        </tr>
      </table>
    </div>
  </div>
</template>

<script>
import NavbarMedico from '../../components/NavbarMedico/NavbarMedico.vue'
import Consultas from '../../services/consultas'

export default {
  data(){
    return{
      solicitacoes: [],
      medicoLogado: {id:1}
    }
  },
  mounted(){
    Consultas.listarPorMedico(this.medicoLogado.id).then(res => {
      this.solicitacoes = res.data.Data.filter(c => {
        return (c.status === false && c.effected === false)
      })
    })
  },
  methods:{
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
    }
  },
  components: {
    NavbarMedico
  }
}

</script>

<style>
#medicoSolicitacoes {
  width: 100vw;
  height: 100vh;
  background-color:  var(--primary);
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
</style>
