import { api } from './config'

export default {
    listarPorPaciente:(clientId) => {
        return api.get(`consultas/client/${clientId}`)
    },

    listarPorMedico:(medicoId) => {
        return api.get(`consultas/medic/${medicoId}`)
    },
    excluir:(id)=>{
        return api.get(`removeconsulta/${id}`)
    },
    atualizar:(consulta)=>{
        api.post('updateconsulta',{consulta})
    }
}