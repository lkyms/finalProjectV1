//
//  MeView.swift
//  finalProjectV1
//
//  Created by Lwj on 2021/11/26.
//

import SwiftUI

struct MeView: View {
    var body: some View {
        NavigationView {
            Form {
                Section (header: Text("Infomation")) {
                    Text("alex")
                }
                NavigationLink {
                    SettingView()
                } label: {
                    Text("Settings")
                }

            }
        }
        .navigationTitle("Me!")
    }
}

struct MeView_Previews: PreviewProvider {
    static var previews: some View {
        MeView()
    }
}
