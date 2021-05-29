const Cicada = artifacts.require("Cicada");

module.exports = function (deployer) {
  deployer.deploy(Cicada, 9000000000000, 10000000000000);
}
