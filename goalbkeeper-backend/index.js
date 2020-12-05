var express = require('express')
var http = require('http')
var mongoose = require('mongoose');

var goalSchema = mongoose.Schema({
    name: {type: String, required: true, unique: true},
    percentComplete: {type: Number, required: true},
    epic: {type: String, required: true},
    createdAt: {type: Date, default: Date.now}
});

var Goal = mongoose.model("Goal", goalSchema);


mongoose.connect('mongodb://localhost/test');


var app = express()

app.get('/save', function (req, res) {
    new Goal({name: 'read', percentComplete: 10, epic: 'general'})
        .save(function (err, goal) {
            if (err) return console.error(err);
            res.send('hello world')
        })
})

app.get('/goals', function (req, res) {
    Goal.find(function (err, goals) {
        if (err) return console.error(err);
        console.log(goals);
        res.json(goals)
    })
})

http.createServer(app).listen(3000)