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
        return api.post('updateconsulta',consulta)
    },
    cadastrar:(consulta)=>{
        return api.post('consulta',consulta)
    }
}