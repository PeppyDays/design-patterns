use super::{adaptee::SpecificTarget, adapter::TargetAdapter, target::{OrdinaryTarget, Target}};

fn call(target: impl Target) {
    println!("'{}'", target.request());
}

fn main() {
    let target = OrdinaryTarget;
    call(target);

    let adaptee = SpecificTarget;
    // unable to call directly
    // call(adaptee);

    let adapter = TargetAdapter::new(adaptee);
    call(adapter);
}
