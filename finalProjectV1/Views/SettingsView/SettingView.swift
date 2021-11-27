//
//  SettingView.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/26.
//

import SwiftUI

struct SettingView: View {
    var body: some View {
        Form{
            NavigationLink {
                ChangeInfoView()
            } label: {
                HStack {
                    Image(systemName: "person")
                    Text("Alex")
                }
            }
        }
        

    }
}

struct SettingView_Previews: PreviewProvider {
    static var previews: some View {
        SettingView()
    }
}
