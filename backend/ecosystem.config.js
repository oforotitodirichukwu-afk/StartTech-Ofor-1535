module.exports = {
  apps : [{
    name: "starttech-backend",
    script: "./main", // Points to your Go binary
    instances: "max",
    exec_mode: "cluster"
  }]
}