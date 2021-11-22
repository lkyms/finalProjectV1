//
//  LoginModel.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/21.
//

import Foundation

struct CheckLogin {
    
    func checkUserNameAndPassword(_ name: String, _ pswd: String) -> Bool {
        if(name == "root" && pswd == "root") {
            return true
        } else {
            return false
        }
    }
    
}
