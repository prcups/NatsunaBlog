import axios from "axios"
import global from "../../../global";

function setIfChecked(t) {
    axios({
        method: "post",
        url: global.CheckUrl
    })
        .then(res => {
            if (t.user !== undefined) t.user = res.data.user
            t.isChecked = res.data.isChecked
        })
}

function logout(t) {
    axios({
        method: "get",
        url: global.LogOutUrl
    })
        .then(res => {
            t.$router.go(0)
        })
}

function toCheck(t) {
    t.$router.push('/admin/check')
}

export {
    logout,
    setIfChecked,
    toCheck
}