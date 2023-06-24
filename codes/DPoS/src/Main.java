import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.util.Base64;

public class Main {

    public static void main(String[] args) {
        // 创建一个具有 4 个难度值的区块链
        Blockchain blockchain = new Blockchain(4);

        // 创建创世区块
        blockchain.createGenesisBlock();

        // 生成密钥对以进行测试
        KeyPair keyPair1 = generateKeyPair();
        KeyPair keyPair2 = generateKeyPair();

        // 获取公钥和私钥
        String publicKey1 = publicKeyToString(keyPair1.getPublic());
        String privateKey1 = privateKeyToString(keyPair1.getPrivate());
        String publicKey2 = publicKeyToString(keyPair2.getPublic());
        String privateKey2 = privateKeyToString(keyPair2.getPrivate());

        // 添加交易
        Blockchain.Transaction transaction1 = new Blockchain.Transaction(publicKey1, publicKey2, 10, privateKey1);
        blockchain.addTransaction(transaction1);

        // 添加投票
        Blockchain.Vote vote1 = new Blockchain.Vote("Candidate A", publicKey1, privateKey1);
        blockchain.addVote(vote1);

        // 添加节点
        Blockchain.Node node1 = new Blockchain.Node("192.168.0.1");
        blockchain.addNode(node1);

        // 创建一个新的区块并添加到区块链
        Blockchain.Block newBlock = new Blockchain.Block(1, System.currentTimeMillis(), blockchain.getTransactionList(),
                blockchain.getVoteList(), blockchain.getLatestBlock().getHash(), 0);
        blockchain.addBlock(newBlock);

        // 验证区块链的合法性
        System.out.println("Blockchain is valid: " + blockchain.validate());
    }

    // 生成密钥对
    private static KeyPair generateKeyPair() {
        try {
            KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
            keyPairGenerator.initialize(2048);
            return keyPairGenerator.generateKeyPair();
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
            return null;
        }
    }

    // 将 PublicKey 转换为字符串
    private static String publicKeyToString(PublicKey publicKey) {
        return Base64.getEncoder().encodeToString(publicKey.getEncoded());
    }

    // 将 PrivateKey 转换为字符串
    private static String privateKeyToString(PrivateKey privateKey) {
        return Base64.getEncoder().encodeToString(privateKey.getEncoded());
    }
}