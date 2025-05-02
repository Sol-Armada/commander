// Ships
import { getMembers } from '@/api'
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { useWebsocketStore } from '@/stores/websocket'
import { WebSocketChannels, WebSocketTypes } from '@/classes/ws'

const wsStore = useWebsocketStore()

const rankToTag = {
    1: "[ADM]",
    2: "[CDR]",
    3: "[LT]",
    4: "[SPC]",
    5: "[TEC]",
    6: "[MBR]",
    7: "[RCT]",
    8: "[GST]",
    99: "[ALY]"
}

const ranktoName = {
    1: "Admiral",
    2: "Commander",
    3: "Lieutenant",
    4: "Specialist",
    5: "Technician",
    6: "Member",
    7: "Recruit",
    8: "Guest",
    99: "Ally"
}

export const useMembersStore = defineStore('members', () => {
    const members = ref([])
    const loading = ref(true)

    function getMember(memberId) {
        const member = members.value.find((member) => member.id == memberId);
        if (member) {
            return member
        }
        return null
    }

    function getNameWithRank(member) {
        return rankToTag[member.rank] + " " + member.name
    }

    function getRankName(rank) {
        return ranktoName[rank]
    }

    function updateMember(member) {
        const index = members.value.findIndex(m => m.id === member.id)
        if (index !== -1) {
            members.value[index] = member
        } else {
            members.value.push(member)
        }
        // const req = new WebSocketResponse(socket, WebSocketTypes.UPDATE, member)
        // req.send()
    }

    function resetMember(memberId) {
        const member = members.value.find(m => m.id === memberId)
        member.shipId = ""
        member.assigned = false
        member.standby = true
        member.squadId = ""

        updateMember(member)
    }

    getMembers()
        .then(res => {
            members.value = res
            loading.value = false
        })
        .catch(err => {
            console.error("Error fetching members:", err)
            loading.value = false
        })

    wsStore.addChannel(WebSocketChannels.MEMBERS, (res) => {
        if (res.type === WebSocketTypes.UPDATED) {
            console.debug("Member updated", res.data);
            updateMember(res.data);
        }
    })

    return {
        members,
        loading,
        getMember,
        updateMember,
        resetMember,
        getNameWithRank,
        getRankName
    }
})
