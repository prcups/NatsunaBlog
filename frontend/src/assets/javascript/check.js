import axios from "axios"
import url from "./url"

function setIfChecked(t) {
    axios({
        method: "post",
        url: url.CheckUrl
    })
        .then(res => {
            if (t.user !== undefined) t.user = res.data.user
            t.isChecked = res.data.isChecked
        })
}

function logout(t) {
    axios({
        method: "get",
        url: url.LogOutUrl
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