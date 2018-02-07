function WelcomeViewModel() {
    var self = this;

    self.isLoad = false;
    self.isConnect = ko.observable(false);
    self.btnText = ko.observable('');

    self.title = C.appName + ' ' + C.appVersion

    self.onLoad = function () {
        if (!(window && window.external && window.external.invoke)) {
            self.isConnect(true);
            self.btnText('go层交互对象错误');
            return;
        }
        self.isLoad = true;

        self.onOpen();
        self.onCheck();
    };

    self.onOpen = function () {
        if (!self.isLoad) return;

        window.external.invoke('open');
    };

    self.onCheck = function () {
        setTimeout(function () {
            window.external.invoke('check');
            self.onCheck();
        }, 1500);
    };

    self.onGithub = function () {
        window.external.invoke('github');
    };

    self.onGitee = function () {
        window.external.invoke('gitee');
    };

    self.onClose = function () {
        window.external.invoke('close');
    };

    self.callJs = function (data) {
        if (!self.isLoad) return;

        if (data === 'connect') {
            self.isConnect(true);
            self.btnText('goman已在浏览器中打开');
        } else if (data === 'disconnect') {
            self.isConnect(false);
            self.btnText('在浏览器中打开goman');
        }
    }
};

var vm = new WelcomeViewModel();
ko.applyBindings(vm);
vm.onLoad();