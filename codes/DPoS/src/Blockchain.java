import java.nio.charset.StandardCharsets;
import java.security.*;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;
import java.util.ArrayList;
import java.util.Base64;
import java.util.List;

public class Blockchain {

    // 区块链列表
    private static List<Blockchain> blockchainList = new ArrayList<Blockchain>();

    // 区块链的难度
    private int difficulty;

    // 交易列表
    private List<Transaction> transactionList = new ArrayList<Transaction>();

    // 投票列表
    private List<Vote> voteList = new ArrayList<Vote>();

    // 节点列表
    private List<Node> nodeList = new ArrayList<Node>();

    // 区块列表
    private List<Block> blockList = new ArrayList<Block>();

    // 构造函数
    public Blockchain(int difficulty) {
        this.difficulty = difficulty;
        blockchainList.add(this);
    }

    // 区块类
    public static class Block {
        private int index;
        private long timestamp;
        private List<Transaction> transactionList;
        private List<Vote> voteList;
        private String previousHash;
        private String hash;
        private int nonce;

        // 构造函数
        public Block(int index, long timestamp, List<Transaction> transactionList, List<Vote> voteList,
                String previousHash, int nonce) {
            this.index = index;
            this.timestamp = timestamp;
            this.transactionList = transactionList;
            this.voteList = voteList;
            this.previousHash = previousHash;
            this.nonce = nonce;
            this.hash = calculateHash();
        }

        // 计算区块的哈希值
        public String calculateHash() {
            String data = index + timestamp + previousHash + transactionList.toString() + voteList.toString() + nonce;
            return HashUtil.sha256(data);
        }

        // 挖矿
        public void mineBlock(int difficulty) {
            String target = new String(new char[difficulty]).replace('\0', '0');
            while (!hash.substring(0, difficulty).equals(target)) {
                nonce++;
                hash = calculateHash();
            }
            System.out.println("Block mined: " + hash);
        }

        // 验证区块的合法性
        public boolean validate() {
            if (!hash.equals(calculateHash())) {
                return false;
            }
            if (!previousHash.equals(getPreviousBlockHash())) {
                return false;
            }
            for (Transaction tx : transactionList) {
                if (!tx.validate()) {
                    return false;
                }
            }
            for (Vote vote : voteList) {
                if (!vote.validate()) {
                    return false;
                }
            }
            return true;
        }

        // 获取上一个区块的哈希值
        public String getPreviousBlockHash() {
            Blockchain blockchain = blockchainList.get(blockchainList.size() - 1);
            List<Block> blockList = blockchain.getBlockList();
            int size = blockList.size();
            if (size == 0) {
                return "";
            }
            return blockList.get(size - 1).hash;
        }

        public int getIndex() {
            return index;
        }

        public long getTimestamp() {
            return timestamp;
        }

        public List<Transaction> getTransactionList() {
            return transactionList;
        }

        public List<Vote> getVoteList() {
            return voteList;
        }

        public String getPreviousHash() {
            return previousHash;
        }

        public String getHash() {
            return hash;
        }

        public int getNonce() {
            return nonce;
        }
    }

    // 交易类
    public static class Transaction {
        private String from;
        private String to;
        private int amount;
        private String signature;

        // 构造函数
        public Transaction(String from, String to, int amount, String privateKey) {
            this.from = from;
            this.to = to;
            this.amount = amount;
            this.signature = SignatureUtil.sign(privateKey, from + to + amount);
        }

        // 验证交易的合法性
        public boolean validate() {
            if (!from.equals("")) {
                if (!SignatureUtil.verify(from, from + to + amount, signature)) {
                    return false;
                }
            }
            return true;
        }

        public String getFrom() {
            return from;
        }

        public String getTo() {
            return to;
        }

        public int getAmount() {
            return amount;
        }

        public String getSignature() {
            return signature;
        }
    }

    // 投票类
    public static class Vote {
        private String candidate;
        private String voter;
        private String signature;

        // 构造函数
        public Vote(String candidate, String voter, String privateKey) {
            this.candidate = candidate;
            this.voter = voter;
            this.signature = SignatureUtil.sign(privateKey, candidate + voter);
        }

        // 验证投票的合法性
        public boolean validate() {
            if (!SignatureUtil.verify(voter, candidate + voter, signature)) {
                return false;
            }
            return true;
        }

        public String getCandidate() {
            return candidate;
        }

        public String getVoter() {
            return voter;
        }

        public String getSignature() {
            return signature;
        }
    }

    // 节点类
    public static class Node {
        private String address;

        // 构造函数
        public Node(String address) {
            this.address = address;
        }

        public String getAddress() {
            return address;
        }
    }

    // 工具类：哈希计算
    public static class HashUtil {
        public static String sha256(String data) {
            try {
                MessageDigest digest = MessageDigest.getInstance("SHA-256");
                byte[] hash = digest.digest(data.getBytes(StandardCharsets.UTF_8));
                return Base64.getEncoder().encodeToString(hash);
            } catch (Exception e) {
                e.printStackTrace();
                return null;
            }
        }
    }

    // 工具类：签名和验证
    public static class SignatureUtil {
        public static String sign(String privateKey, String data) {
            try {
                PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(Base64.getDecoder().decode(privateKey));
                KeyFactory keyFactory = KeyFactory.getInstance("RSA");
                PrivateKey key = keyFactory.generatePrivate(keySpec);
                Signature signature = Signature.getInstance("SHA256withRSA");
                signature.initSign(key);
                signature.update(data.getBytes(StandardCharsets.UTF_8));
                byte[] signBytes = signature.sign();
                return Base64.getEncoder().encodeToString(signBytes);
            } catch (NoSuchAlgorithmException | InvalidKeySpecException | InvalidKeyException | SignatureException e) {
                e.printStackTrace();
                return null;
            }
        }

        public static boolean verify(String publicKey, String data, String signatureString) {
            try {
                X509EncodedKeySpec keySpec = new X509EncodedKeySpec(Base64.getDecoder().decode(publicKey));
                KeyFactory keyFactory = KeyFactory.getInstance("RSA");
                PublicKey key = keyFactory.generatePublic(keySpec);
                Signature signature = Signature.getInstance("SHA256withRSA");
                signature.initVerify(key);
                signature.update(data.getBytes(StandardCharsets.UTF_8));
                byte[] signatureBytes = Base64.getDecoder().decode(signatureString);
                return signature.verify(signatureBytes);
            } catch (NoSuchAlgorithmException | InvalidKeySpecException | InvalidKeyException | SignatureException e) {
                e.printStackTrace();
                return false;
            }
        }
    }

    public int getDifficulty() {
        return difficulty;
    }

    public List<Transaction> getTransactionList() {
        return transactionList;
    }

    public List<Vote> getVoteList() {
        return voteList;
    }

    public List<Node> getNodeList() {
        return nodeList;
    }

    public List<Block> getBlockList() {
        return blockList;
    }

    // 创建创世块
    public void createGenesisBlock() {
        List<Transaction> transactions = new ArrayList<Transaction>();
        List<Vote> votes = new ArrayList<Vote>();
        String previousHash = "";
        int nonce = 0;
        Block genesisBlock = new Block(0, System.currentTimeMillis(), transactions, votes, previousHash, nonce);
        blockList.add(genesisBlock);
    }

    // 添加交易
    public void addTransaction(Transaction transaction) {
        transactionList.add(transaction);
    }

    // 添加投票
    public void addVote(Vote vote) {
        voteList.add(vote);
    }

    // 添加节点
    public void addNode(Node node) {
        nodeList.add(node);
    }

    // 获取最新的区块
    public Block getLatestBlock() {
        List<Block> blockList = getBlockList();
        return blockList.get(blockList.size() - 1);
    }

    // 添加区块
    public void addBlock(Block block) {
        if (block.validate() && block.getPreviousHash().equals(getLatestBlock().hash)) {
            block.mineBlock(difficulty);
            blockList.add(block);
        }
    }

    // 验证区块链的合法性
    public boolean validate() {
        List<Block> blockList = getBlockList();
        for (int i = 1; i < blockList.size(); i++) {
            Block currentBlock = blockList.get(i);
            Block previousBlock = blockList.get(i - 1);
            if (!currentBlock.validate() || !currentBlock.getPreviousHash().equals(previousBlock.hash)) {
                return false;
            }
        }
        return true;
    }
}