package path

var BaseUrl = "/api"
var baseVersion = "/v1"
var id = "/:id"

var Soldier = baseVersion + "/soldier"
var SoldierById = Soldier + id

var Magazine = baseVersion + "/magazine"
var MagazineLoad = Magazine + "/load"