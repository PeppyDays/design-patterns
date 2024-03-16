use crate::structure::facade::wallet_facade::WalletFacade;

fn main() -> Result<(), String> {
    let mut wallet = WalletFacade::new("abc".into(), 1234);
    println!();

    wallet.add_money_to_wallet("abc", 1234, 10)?;
    wallet.deduct_money_from_wallet("abc", 1234, 5)?;

    Ok(())
}
