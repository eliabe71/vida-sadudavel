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
                <h5 class="card-title">Solicitação</h5>
                <p class="card-text"><b>Paciente:</b>{{s.clientName}} </p>
                <p class="card-text"><b>Data: </b>{{formatDate(s.day)}}</p>
                <p class="card-text"><b>Horário:</b> {{formatHour(s.hourInit)}}</p>
              </div>
              <div class="cart-footer">
                <button class="btn-success" data-bs-toggle="modal" :data-bs-target="'#confModalexcluir'+s.id" >Aceitar</button>
              </div>
            </div>
            <div class="modal fade" :id="'confModalexcluir'+s.id" tabindex="-1" aria-labelledby="confimationModalLabel" aria-hidden="true">
              <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                  <div class="modal-header">
                    <h5 class="modal-title" id="confimationModalLabel">Excluir consulta</h5>
                    <button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
                  </div>
                  <div class="modal-body">
                    <p>Deseja realmente aceitar a solicitação de consulta?</p>
                  </div>
                  <div class="modal-footer">
                    <button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Cancelar</button>
                    <button type="button" v-on:click="aceitarConsulta(s.id)" data-bs-dismiss="modal" class="btn btn-success">Confirmar</button>
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
        res = res[0]+'-'+res[1]+'-'+res[1]
        return res
      },
      formatHour(hour){
        let res = hour.split('T')
        res = res[1].split(':')
        res = res[0]+':'+res[1]
        return res
      },
      aceitarConsulta(id){
        let cons = this.solicitacoes.filter(s => {
          return (s.id === id)
        })

        cons[0].status = true
        cons[0].day = this.formatDateJson(cons[0].day)
        cons[0].hourInit = this.formatHour(cons[0].hourInit)
        cons[0].hourEnd = this.formatHour(cons[0].hourEnd)
        console.log(cons[0])

        Consultas.atualizar(cons[0]).then(res => {
          if(res.status === 200){
            this.$router.go()
            alert("Consulta aceita com sucesso!")
          }
          else
            alert("Erro ao aceitar consulta!")
        })
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
